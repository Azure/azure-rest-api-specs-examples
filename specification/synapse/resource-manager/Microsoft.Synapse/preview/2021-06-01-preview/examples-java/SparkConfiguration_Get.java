/** Samples for SparkConfiguration Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/SparkConfiguration_Get.json
     */
    /**
     * Sample code: Get SparkConfiguration by name.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getSparkConfigurationByName(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sparkConfigurations()
            .getWithResponse(
                "exampleResourceGroup",
                "exampleSparkConfigurationName",
                "exampleWorkspace",
                com.azure.core.util.Context.NONE);
    }
}
