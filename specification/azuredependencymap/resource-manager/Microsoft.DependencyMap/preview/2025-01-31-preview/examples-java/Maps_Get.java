
/**
 * Samples for Maps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-31-preview/Maps_Get.json
     */
    /**
     * Sample code: Maps_Get - generated by [MaximumSet] rule.
     * 
     * @param manager Entry point to DependencyMapManager.
     */
    public static void
        mapsGetGeneratedByMaximumSetRule(com.azure.resourcemanager.dependencymap.DependencyMapManager manager) {
        manager.maps().getByResourceGroupWithResponse("rgdependencyMap", "mapsTest1", com.azure.core.util.Context.NONE);
    }
}
