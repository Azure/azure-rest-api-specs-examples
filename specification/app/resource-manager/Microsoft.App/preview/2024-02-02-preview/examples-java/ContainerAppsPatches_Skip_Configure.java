
import com.azure.resourcemanager.appcontainers.models.PatchSkipConfig;

/**
 * Samples for ContainerAppsPatches SkipConfigure.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ContainerAppsPatches_Skip_Configure.json
     */
    /**
     * Sample code: ContainerAppsPatches_Skip_Configure_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        containerAppsPatchesSkipConfigure0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsPatches().skipConfigure("rg", "test-app", "testPatch-25fe4b",
            new PatchSkipConfig().withSkip(true), com.azure.core.util.Context.NONE);
    }
}
