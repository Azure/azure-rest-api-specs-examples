
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.appcontainers.models.WorkflowArtifacts;
import java.io.IOException;

/**
 * Samples for LogicApps DeployWorkflowArtifacts.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * LogicApps_PostDeployWorkflowArtifacts.json
     */
    /**
     * Sample code: Deploys workflow artifacts.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deploysWorkflowArtifacts(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager)
        throws IOException {
        manager.logicApps().deployWorkflowArtifactsWithResponse("testrg123", "testapp2", "testapp2",
            new WorkflowArtifacts()
                .withAppSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"eventHub_connectionString\":\"Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=EXAMPLE1a2b3c4d5e6fEXAMPLE=\"}",
                    Object.class, SerializerEncoding.JSON))
                .withFiles(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"connections.json\":{\"managedApiConnections\":{},\"serviceProviderConnections\":{\"eventHub\":{\"displayName\":\"example1\",\"parameterValues\":{\"connectionString\":\"@appsetting('eventHub_connectionString')\"},\"serviceProvider\":{\"id\":\"/serviceProviders/eventHub\"}}}},\"test1/workflow.json\":{\"definition\":{\"$schema\":\"https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#\",\"actions\":{},\"contentVersion\":\"1.0.0.0\",\"outputs\":{},\"triggers\":{\"When_events_are_available_in_Event_hub\":{\"type\":\"ServiceProvider\",\"inputs\":{\"parameters\":{\"eventHubName\":\"test123\"},\"serviceProviderConfiguration\":{\"operationId\":\"receiveEvents\",\"connectionName\":\"eventHub\",\"serviceProviderId\":\"/serviceProviders/eventHub\"}},\"splitOn\":\"@triggerOutputs()?['body']\"}}},\"kind\":\"Stateful\"}}",
                    Object.class, SerializerEncoding.JSON)),
            com.azure.core.util.Context.NONE);
    }
}
