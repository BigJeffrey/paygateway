package postgresqldao

import (
	"log"
	"paygateway/models"

	"github.com/google/uuid"
)

func (m *PostgreSql) GetCards() ([]models.CreditCardData, error) {
	var creditCard models.CreditCardData
	var creditCards []models.CreditCardData
	sqlStatement := `SELECT * FROM creditcarddata`

	rows, err := m.client.Query(sqlStatement)
	if err != nil {
		return []models.CreditCardData{}, err
	}

	for rows.Next() {
		err := rows.Scan(&creditCard.ID, &creditCard.CardNumber, &creditCard.ExpireMonthDay, &creditCard.Cvv, &creditCard.Amount, &creditCard.Currency)
		if err != nil {
			return nil, err
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards, nil
}

func (m *PostgreSql) UpdateCard(cardData models.CreditCardData) error {
	sqlStatement := `UPDATE creditcarddata SET expiremothdate=$1, cvv=$2, amount=$3, currency=$4 WHERE cardnumber=$5`
	_, err := m.client.Exec(sqlStatement, cardData.ExpireMonthDay, cardData.Cvv, cardData.Amount, cardData.Currency, cardData.CardNumber)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgreSql) GetMerchantByLogin(login string) (*models.Merchant, error) {
	var merchant models.Merchant
	sqlStatement := `SELECT * FROM merchants WHERE username=$1`
	err := m.client.QueryRow(sqlStatement, login).Scan(&merchant.UUID, &merchant.Username, &merchant.Password)
	if err != nil {
		log.Println(err)
		return &models.Merchant{}, err

	}

	return &merchant, nil
}

func (m *PostgreSql) GetTransaction(transactionId *uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	sqlStatement := `SELECT * FROM transactions WHERE id=$1`

	err := m.client.QueryRow(sqlStatement, transactionId).Scan(&transaction.UUID, &transaction.Status, &transaction.PaymentMethod, &transaction.CardNumber, &transaction.ExpireMonthDay, &transaction.Cvv, &transaction.Amount, &transaction.Currency, &transaction.Spent, &transaction.ProviderTransactionId)
	if err != nil {
		log.Println(err, transactionId)
		return &models.Transaction{}, err
	}

	return &transaction, nil
}

func (m *PostgreSql) CreateTransaction(transaction *models.Transaction) (uuid.UUID, error) {
	var id uuid.UUID
	sqlStatement := `INSERT INTO transactions (status, paymentmethod, cardnumber, expiremonthday, cvv, amount, currency, spent, providertransactionid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err := m.client.QueryRow(sqlStatement, transaction.Status, transaction.PaymentMethod, transaction.CardNumber, transaction.ExpireMonthDay, transaction.Cvv, transaction.Amount, transaction.Currency, transaction.Spent, transaction.ProviderTransactionId).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *PostgreSql) UpdateTransaction(transaction *models.Transaction) error {
	sqlStatement := `UPDATE transactions SET status=$1, paymentmethod=$2, cardnumber=$3, expiremonthday=$4, cvv=$5, amount=$6, currency=$7, spent=$8, providertransactionid=$9 WHERE id=$10`
	_, err := m.client.Exec(sqlStatement, transaction.Status, transaction.PaymentMethod, transaction.CardNumber, transaction.ExpireMonthDay, transaction.Cvv, transaction.Amount, transaction.Currency, transaction.Spent, transaction.ProviderTransactionId, transaction.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgreSql) DeleteTransaction(transaction models.Transaction) error {
	sqlStatement := `DELETE FROM transactions WHERE id=$1`
	_, err := m.client.Exec(sqlStatement, transaction.UUID)
	if err != nil {
		return err
	}
	return nil
}
