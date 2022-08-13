import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for DscCompilationJobStream ListByJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/compilationJobStreamList.json
     */
    /**
     * Sample code: List DSC Compilation job streams.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listDSCCompilationJobStreams(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscCompilationJobStreams()
            .listByJobWithResponse(
                "rg", "myAutomationAccount33", UUID.fromString("836d4e06-2d88-46b4-8500-7febd4906838"), Context.NONE);
    }
}
