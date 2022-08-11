import com.azure.core.util.Context;

/** Samples for BestPractices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getBestPractice.json
     */
    /**
     * Sample code: Get an Automanage best practice.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAnAutomanageBestPractice(com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.bestPractices().getWithResponse("azureBestPracticesProduction", Context.NONE);
    }
}
