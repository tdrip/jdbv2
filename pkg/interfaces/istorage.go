package interfaces

type IStorage interface {
	ReadOnly() bool
	Exit() chan error
	Run(items map[string]IKeyedItem) (map[string]IKeyedItem, error)
}
