
/**
 * Samples for ResourceHealthMetadata ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListResourceHealthMetadataByResourceGroup.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listResourceHealthMetadataForAResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceHealthMetadatas().listByResourceGroup("Default-Web-NorthCentralUS",
            com.azure.core.util.Context.NONE);
    }
}
