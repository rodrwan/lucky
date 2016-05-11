import random

txt = open("training_data2.txt")
context = txt.read()
samples = context.split("\n")
samples.pop(-1)
maxSamples = len(samples)

print maxSamples
randomSamples = random.sample(range(0, maxSamples), int(maxSamples * 0.1))
txt.close()
print len(randomSamples)

sampleTarget = open("test.txt", "w")
for key in randomSamples:
    sampleTarget.write(samples[key] + "\n")
sampleTarget.close()

randomSamples = sorted(randomSamples, reverse=True)
for key in randomSamples:
    samples.pop(key)

sampleTrain = open("train.txt", "w")
maxSamples = len(samples)
print maxSamples
for key in range(0, maxSamples):
    sampleTrain.write(samples[key] + "\n")
sampleTrain.close()
print "job done"
