package messagesService

type MessageService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) PatchMessage(id int, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

func (s *MessageService) DeleteMessage(id int) error {
	return s.repo.DeleteMessageByID(id)
}
