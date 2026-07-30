
/**
 * Samples for PreparedImageSpecifications GetVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/PreparedImageSpecifications_GetVersion.json
     */
    /**
     * Sample code: PreparedImageSpecifications_GetVersion.
     * 
     * @param manager Entry point to ContainerServicePreparedImageSpecificationManager.
     */
    public static void preparedImageSpecificationsGetVersion(
        com.azure.resourcemanager.containerservicepreparedimgspec.ContainerServicePreparedImageSpecificationManager manager) {
        manager.preparedImageSpecifications().getVersionWithResponse("rg1", "my-prepared-image-specification",
            "20250101-abcd1234", com.azure.core.util.Context.NONE);
    }
}
