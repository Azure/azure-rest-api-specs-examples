
import com.azure.resourcemanager.sql.fluent.models.JobPrivateEndpointInner;

/**
 * Samples for JobPrivateEndpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobPrivateEndpoint.json
     */
    /**
     * Sample code: Create a private endpoint.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createAPrivateEndpoint(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobPrivateEndpoints().createOrUpdate("group1", "server1", "agent1", "endpoint1",
            new JobPrivateEndpointInner().withTargetServerAzureResourceId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
            com.azure.core.util.Context.NONE);
    }
}
