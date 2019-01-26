// Worker is abstraction for in-server-network authentication
// which uses RabbitMQ as bridge between applications in server-network.
// Worker has next work structure:
// authRequestsSource -> broker -> worker -> broker -> authRequestDestination
// In current moment as broker is used RabbitMQ.
package worker
