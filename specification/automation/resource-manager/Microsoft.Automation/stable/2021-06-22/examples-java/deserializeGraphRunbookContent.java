import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.fluent.models.GraphicalRunbookContentInner;
import com.azure.resourcemanager.automation.models.GraphRunbookType;
import com.azure.resourcemanager.automation.models.RawGraphicalRunbookContent;

/** Samples for ResourceProvider ConvertGraphRunbookContent. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deserializeGraphRunbookContent.json
     */
    /**
     * Sample code: Get Graphical runbook content from raw content.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getGraphicalRunbookContentFromRawContent(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .resourceProviders()
            .convertGraphRunbookContentWithResponse(
                "rg",
                "MyAutomationAccount",
                new GraphicalRunbookContentInner()
                    .withRawContent(
                        new RawGraphicalRunbookContent()
                            .withSchemaVersion("1.10")
                            .withRunbookDefinition(
                                "AAEAAADAQAAAAAAAAAMAgAAAGJPcmNoZXN0cmF0b3IuR3JhcGhSdW5ib29rLk1vZGVsLCBWZXJzaW9uPTcuMy4wLjAsIEN1bHR....")
                            .withRunbookType(GraphRunbookType.GRAPH_POWER_SHELL)),
                Context.NONE);
    }
}
