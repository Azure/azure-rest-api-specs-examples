
import com.azure.resourcemanager.devopsinfrastructure.models.Pool;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/UpdatePool.json
     */
    /**
     * Sample code: Pools_Update.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void poolsUpdate(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        Pool resource
            = manager.pools().getByResourceGroupWithResponse("rg", "pool", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
