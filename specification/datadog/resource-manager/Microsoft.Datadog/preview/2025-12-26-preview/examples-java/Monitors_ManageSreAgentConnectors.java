
import com.azure.resourcemanager.datadog.models.ConnectorAction;
import com.azure.resourcemanager.datadog.models.SreAgentConfiguration;
import com.azure.resourcemanager.datadog.models.SreAgentConnectorRequest;
import java.util.Arrays;

/**
 * Samples for Monitors ManageSreAgentConnectors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_ManageSreAgentConnectors.json
     */
    /**
     * Sample code: Monitors_ManageSreAgentConnectors.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void
        monitorsManageSreAgentConnectors(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().manageSreAgentConnectorsWithResponse("myResourceGroup", "myMonitor",
            new SreAgentConnectorRequest().withMcpConnectorResourceIdList(Arrays.asList(
                new SreAgentConfiguration().withMcpConnectorResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.App/agents/sreAgent/connectors/myMcpConnector1"),
                new SreAgentConfiguration().withMcpConnectorResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.App/agents/otherSreAgent/connectors/myMcpConnector2")))
                .withAction(ConnectorAction.ADD),
            com.azure.core.util.Context.NONE);
    }
}
