import com.azure.core.util.Context;

/** Samples for SparkConfigurationsOperation ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/SparkConfigurations_ListByWorkspace.json
     */
    /**
     * Sample code: List sparkConfigurations in a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listSparkConfigurationsInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sparkConfigurationsOperations()
            .listByWorkspace("exampleResourceGroup", "exampleWorkspace", Context.NONE);
    }
}
