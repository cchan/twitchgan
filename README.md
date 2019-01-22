TwitchGAN
=========
[@cchan](https://github.com/cchan), [@RobertRiachi](https://github.com/robertriachi)

Generating 28x28 Twitch emotes with a Wasserstein GAN!

We used Colab to train the model, using their free K80 gpu. #ThanksGoogle

Next steps might be to initialize with Inception weights or even augment the data with ImageNet samples to help pull it toward more meaningful-looking objects.

Examples
--------

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

The training stopped in the middle of Epoch 42, due to Colab timing out the instance.


Files
-----

- downloader.go downloads all Twitch emotes in parallel. It's quite fast.
- png2hdf5.ipynb transforms the ones we want (28x28 RGBA, which is the most common dimensions) into an HDF5 file.
- dataset-first-50k is the first 50,000 28x28 RGBA Twitch emotes, which is what we're training on.
- gan.ipynb contains the actual training code.
