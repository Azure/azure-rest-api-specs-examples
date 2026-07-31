
/**
 * Samples for PreparedImageSpecifications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/PreparedImageSpecifications_ListBySubscription.json
     */
    /**
     * Sample code: PreparedImageSpecifications_ListBySubscription.
     * 
     * @param manager Entry point to ContainerServicePreparedImageSpecificationManager.
     */
    public static void preparedImageSpecificationsListBySubscription(
        com.azure.resourcemanager.containerservicepreparedimgspec.ContainerServicePreparedImageSpecificationManager manager) {
        manager.preparedImageSpecifications().list(com.azure.core.util.Context.NONE);
    }
}
