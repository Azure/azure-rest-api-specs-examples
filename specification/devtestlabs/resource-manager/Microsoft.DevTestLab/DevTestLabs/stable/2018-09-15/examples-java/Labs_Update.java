
import com.azure.resourcemanager.devtestlabs.models.Lab;

/**
 * Samples for Labs Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Labs_Update.json
     */
    /**
     * Sample code: Labs_Update.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        Lab resource = manager.labs()
            .getByResourceGroupWithResponse("resourceGroupName", "{labName}", null, com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
