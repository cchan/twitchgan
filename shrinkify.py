with h5py.File("dataset.hdf5", 'r') as f:
  with h5py.File("dataset-first-50k.hdf5", 'x') as s:
    dset = s.create_dataset("twitch_emotes", (0, 28, 28, 4), dtype='uint8', chunks=(1000,28,28,4), maxshape=(None,28,28,4),compression='gzip',shuffle=True)
    dset.resize(50000, axis=0)
    dset[:]=f['twitch_emotes'][:50000]
    print(dset.shape)
