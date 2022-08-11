import com.azure.core.util.Context;

/** Samples for BestPractices ListByTenant. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listBestPracticesByTenant.json
     */
    /**
     * Sample code: List Automanage bestPractices.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listAutomanageBestPractices(com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.bestPractices().listByTenant(Context.NONE);
    }
}
