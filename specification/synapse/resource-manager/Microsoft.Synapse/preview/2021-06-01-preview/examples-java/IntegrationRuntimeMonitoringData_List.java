import com.azure.core.util.Context;

/** Samples for IntegrationRuntimeMonitoringData List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeMonitoringData_List.json
     */
    /**
     * Sample code: Get monitoring data.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getMonitoringData(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeMonitoringDatas()
            .listWithResponse("exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", Context.NONE);
    }
}
