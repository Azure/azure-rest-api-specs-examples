import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.FlowLogInner;
import com.azure.resourcemanager.network.models.FlowLogFormatParameters;
import com.azure.resourcemanager.network.models.FlowLogFormatType;

/** Samples for FlowLogs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkWatcherFlowLogCreate.json
     */
    /**
     * Sample code: Create or update flow log.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateFlowLog(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFlowLogs()
            .createOrUpdate(
                "rg1",
                "nw1",
                "fl",
                new FlowLogInner()
                    .withLocation("centraluseuap")
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg")
                    .withStorageId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe")
                    .withEnabled(true)
                    .withFormat(new FlowLogFormatParameters().withType(FlowLogFormatType.JSON).withVersion(1)),
                Context.NONE);
    }
}
