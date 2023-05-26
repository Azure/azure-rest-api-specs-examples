import com.azure.resourcemanager.resources.fluent.models.ResourceGroupInner;

/** Samples for ResourceGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/CreateResourceGroup.json
     */
    /**
     * Sample code: Create or update a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getResourceGroups()
            .createOrUpdateWithResponse(
                "my-resource-group", new ResourceGroupInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
