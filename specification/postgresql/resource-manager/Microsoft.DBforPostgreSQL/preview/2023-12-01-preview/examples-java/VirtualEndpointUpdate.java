
import com.azure.resourcemanager.postgresqlflexibleserver.models.VirtualEndpointResource;
import com.azure.resourcemanager.postgresqlflexibleserver.models.VirtualEndpointType;
import java.util.Arrays;

/**
 * Samples for VirtualEndpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/
     * VirtualEndpointUpdate.json
     */
    /**
     * Sample code: Update a virtual endpoint for a server to update the.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateAVirtualEndpointForAServerToUpdateThe(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        VirtualEndpointResource resource = manager.virtualEndpoints()
            .getWithResponse("testrg", "pgtestsvc4", "pgVirtualEndpoint1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withEndpointType(VirtualEndpointType.READ_WRITE).withMembers(Arrays.asList("testReplica1"))
            .apply();
    }
}
