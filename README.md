TwitchGAN
=========
[@cchan](https://github.com/cchan), [@RobertRiachi](https://github.com/robertriachi)

Generating Twitch emotes with a Wasserstein GAN!

We used Colab to train the model, using their free K80 gpu. #ThanksGoogle

### Epoch 0

![Epoch 0 Sample 3](captures/0_3.png?raw=true)

It's already learned what edges are and how to deal with transparency.

### Epoch 2

![Epoch 2 Sample 2](captures/2_2.png?raw=true)

### Epoch 5

![Epoch 5 Sample 1](captures/5_1.png?raw=true)

It seems to be resolving things into objects, ish. An apple core?

### Epoch 13

![Epoch 13 Sample 2](captures/13_2.png?raw=true)

Not much has improved; this seems converged.

### Epoch 31

![Epoch 31 Sample 3](captures/31_3.png?raw=true)

A firefox?

### Epoch 41

![Epoch 41 Sample 1](captures/41_1.png?raw=true)

A pokemon?

The training stopped in the middle of Epoch 42, due to Colab shutting down the instance.


