
/**
 * Samples for SubnetServiceAssociationLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/
     * SubnetServiceAssociationLinkDelete.json
     */
    /**
     * Sample code: SubnetServiceAssociationLinkDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void subnetServiceAssociationLinkDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getSubnetServiceAssociationLinks().delete("demo", "demo2",
            "demo3", com.azure.core.util.Context.NONE);
    }
}
