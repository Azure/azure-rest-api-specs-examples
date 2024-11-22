
/**
 * Samples for ImageVersions ListByImage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/ImageVersions_ListByImage.json
     */
    /**
     * Sample code: ImageVersions_ListByImage.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        imageVersionsListByImage(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.imageVersions().listByImage("my-resource-group", "windows-2022", com.azure.core.util.Context.NONE);
    }
}
