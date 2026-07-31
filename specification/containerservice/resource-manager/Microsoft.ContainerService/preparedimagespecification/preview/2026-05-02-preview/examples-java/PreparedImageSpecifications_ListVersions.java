
/**
 * Samples for PreparedImageSpecifications ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/PreparedImageSpecifications_ListVersions.json
     */
    /**
     * Sample code: PreparedImageSpecifications_ListVersions.
     * 
     * @param manager Entry point to ContainerServicePreparedImageSpecificationManager.
     */
    public static void preparedImageSpecificationsListVersions(
        com.azure.resourcemanager.containerservicepreparedimgspec.ContainerServicePreparedImageSpecificationManager manager) {
        manager.preparedImageSpecifications().listVersions("rg1", "my-prepared-image-specification",
            com.azure.core.util.Context.NONE);
    }
}
