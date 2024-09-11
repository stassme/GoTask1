package messagesService

type MessageService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(m Message) (Message, error) {
	return s.repo.CreateMessage(m)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) UpdateMessageByID(id int, m Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, m)
}

func (s *MessageService) DeleteMessageByID(id int) error {
	return s.repo.DeleteMessageByID(id)
}
