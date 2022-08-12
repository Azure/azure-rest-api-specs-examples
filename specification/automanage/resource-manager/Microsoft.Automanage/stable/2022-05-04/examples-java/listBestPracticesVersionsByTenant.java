import com.azure.core.util.Context;

/** Samples for BestPracticesVersions ListByTenant. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listBestPracticesVersionsByTenant.json
     */
    /**
     * Sample code: List Automanage best practices versions.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listAutomanageBestPracticesVersions(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.bestPracticesVersions().listByTenant("azureBestPracticesProduction", Context.NONE);
    }
}
