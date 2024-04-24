
/**
 * Samples for SolutionSelfHelp Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/SelfHelpSolution_Get.json
     */
    /**
     * Sample code: Solution_Get.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void solutionGet(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.solutionSelfHelps().getWithResponse("SolutionId1", com.azure.core.util.Context.NONE);
    }
}
