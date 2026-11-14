
/**
 * Samples for SubnetServiceAssociationLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SubnetServiceAssociationLinkDelete.json
     */
    /**
     * Sample code: SubnetServiceAssociationLinkDelete.
     * 
     * @param manager Entry point to ContainerInstanceManager.
     */
    public static void subnetServiceAssociationLinkDelete(
        com.azure.resourcemanager.containerinstance.ContainerInstanceManager manager) {
        manager.serviceClient().getSubnetServiceAssociationLinks().delete("demo", "demo2", "demo3",
            com.azure.core.util.Context.NONE);
    }
}
