package main

// Bad structure
type BadReport interface {
	open()
	save()
	BadMongoDB
}

type BadMongoDB interface {
	insert()
	update()
	delete()
}

// Correct structure
type Report interface {
	open()
	save()
	DatabaseProvider
}

type DatabaseProvider interface {
	insert()
	update()
	delete()
}

type MongoDB interface {
	insert()
	update()
	delete()
}

type MySQL interface {
	insert()
	update()
	delete()
}
