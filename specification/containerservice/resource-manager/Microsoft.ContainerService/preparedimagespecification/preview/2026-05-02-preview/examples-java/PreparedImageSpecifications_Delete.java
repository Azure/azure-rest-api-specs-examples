
/**
 * Samples for PreparedImageSpecifications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/PreparedImageSpecifications_Delete.json
     */
    /**
     * Sample code: PreparedImageSpecifications_Delete.
     * 
     * @param manager Entry point to ContainerServicePreparedImageSpecificationManager.
     */
    public static void preparedImageSpecificationsDelete(
        com.azure.resourcemanager.containerservicepreparedimgspec.ContainerServicePreparedImageSpecificationManager manager) {
        manager.preparedImageSpecifications().delete("rg1", "my-prepared-image-specification", null,
            com.azure.core.util.Context.NONE);
    }
}
