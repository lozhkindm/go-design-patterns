package main

type NeuronObject interface {
	Collect() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Collect() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(neuron *Neuron) {
	n.Out = append(n.Out, neuron)
	neuron.In = append(neuron.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Collect() []*Neuron {
	neurons := make([]*Neuron, len(n.Neurons))
	for i := range n.Neurons {
		neurons = append(neurons, &n.Neurons[i])
	}
	return neurons
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{Neurons: make([]Neuron, count)}
}

func Connect(left, right NeuronObject) {
	for _, l := range left.Collect() {
		for _, r := range right.Collect() {
			l.ConnectTo(r)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
}
