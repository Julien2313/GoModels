Comment : I wrote what I understood for every method, things could be wrong. I Always try to explain why, but I think things are getting quickly hard... So, if you have any advive :)

# OLS
## 1. Dataset

For this method, we need a linear problem, because the solution has this shape : a.x+b, and this is a linear function.

## 2. what are we looking for ?

We need to find the line which will minimise the squared errors. For this case, the errors are the squared difference of Y between the dataset and the model.

$$χ^2=\sum_{i=1}^N{(y_i-p_i)^2}$$(1)
Where `'y'` is the calculated value according to our model and `'y'` the value from the dataset

## 3. Let's try things

I think that one of the good start could be to calculate the mean of y and x of the dataset, and because all the data are equaly importants, this would be a kind of barycenter.

$$\bar x={\sum_{i=1}^N{x_i} \over N}$$(2)
$$\bar y={\sum_{i=1}^N{y_i} \over N}$$(3)

So, we need to find a line that cross this point, let's find out the equation !