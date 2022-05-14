import numpy as np
import matplotlib.pyplot as plt

incomes = np.random.normal(28000, 15000, 10000)
print(np.mean(incomes))

plt.hist(incomes, 50)
plt.show()