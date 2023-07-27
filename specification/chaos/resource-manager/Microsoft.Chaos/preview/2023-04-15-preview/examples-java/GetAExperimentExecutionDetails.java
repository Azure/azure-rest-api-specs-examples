/** Samples for Experiments GetExecutionDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetAExperimentExecutionDetails.json
     */
    /**
     * Sample code: Get experiment execution details.
     *
     * @param manager Entry point to ChaosManager.
     */
    public static void getExperimentExecutionDetails(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager
            .experiments()
            .getExecutionDetailsWithResponse(
                "exampleRG",
                "exampleExperiment",
                "f24500ad-744e-4a26-864b-b76199eac333",
                com.azure.core.util.Context.NONE);
    }
}
