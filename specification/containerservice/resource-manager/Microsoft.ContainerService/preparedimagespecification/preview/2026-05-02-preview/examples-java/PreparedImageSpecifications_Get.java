
/**
 * Samples for PreparedImageSpecifications GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/PreparedImageSpecifications_Get.json
     */
    /**
     * Sample code: PreparedImageSpecifications_Get.
     * 
     * @param manager Entry point to ContainerServicePreparedImageSpecificationManager.
     */
    public static void preparedImageSpecificationsGet(
        com.azure.resourcemanager.containerservicepreparedimgspec.ContainerServicePreparedImageSpecificationManager manager) {
        manager.preparedImageSpecifications().getByResourceGroupWithResponse("rg1", "my-prepared-image-specification",
            com.azure.core.util.Context.NONE);
    }
}
