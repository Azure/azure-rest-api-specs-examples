import com.azure.core.util.Context;

/** Samples for BestPracticesVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getBestPracticeVersion.json
     */
    /**
     * Sample code: Get an Automanage best practice version.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAnAutomanageBestPracticeVersion(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.bestPracticesVersions().getWithResponse("azureBestPracticesProduction", "version1", Context.NONE);
    }
}
