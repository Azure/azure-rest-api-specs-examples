
/**
 * Samples for SolutionSelfHelp Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01-preview/SelfHelpSolution_Get.json
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
