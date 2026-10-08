
import com.azure.resourcemanager.resources.fluent.models.ResourceGroupInner;

/**
 * Samples for ResourceGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/CreateResourceGroup.json
     */
    /**
     * Sample code: Create or update a resource group.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void createOrUpdateAResourceGroup(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getResourceGroups().createOrUpdateWithResponse("my-resource-group",
            new ResourceGroupInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
