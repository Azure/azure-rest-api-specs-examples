
/**
 * Samples for PreparedImageSpecifications ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/PreparedImageSpecifications_ListByResourceGroup.json
     */
    /**
     * Sample code: PreparedImageSpecifications_ListByResourceGroup.
     * 
     * @param manager Entry point to ContainerServicePreparedImageSpecificationManager.
     */
    public static void preparedImageSpecificationsListByResourceGroup(
        com.azure.resourcemanager.containerservicepreparedimgspec.ContainerServicePreparedImageSpecificationManager manager) {
        manager.preparedImageSpecifications().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
