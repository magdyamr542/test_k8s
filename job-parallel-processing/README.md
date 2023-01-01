# Parallel processing using jobs and a work queue

1. A RabbitMQ queue is started and filled with some tasks
2. A job is started to process the tasks from the queue (can be emails or file to be processed etc.)

# Running

1. Push the required images to the kind control plane with `make prepare`
2. Run the broker with `make broker`
3. Run the tester to create a queue and add some tasks with `make connect-tester`
   - follow the details [here](https://kubernetes.io/docs/tasks/job/coarse-parallel-processing-work-queue/)
   - the queue name should be `send-email` unlike the tutorial
4. Run the job to process some tasks from the queue
