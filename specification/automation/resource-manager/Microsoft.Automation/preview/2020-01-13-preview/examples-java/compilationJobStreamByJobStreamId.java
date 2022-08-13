import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for DscCompilationJob GetStream. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/compilationJobStreamByJobStreamId.json
     */
    /**
     * Sample code: Get a DSC Compilation job stream by job stream id.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getADSCCompilationJobStreamByJobStreamId(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .dscCompilationJobs()
            .getStreamWithResponse(
                "rg",
                "myAutomationAccount33",
                UUID.fromString("836d4e06-2d88-46b4-8500-7febd4906838"),
                "836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008",
                Context.NONE);
    }
}
