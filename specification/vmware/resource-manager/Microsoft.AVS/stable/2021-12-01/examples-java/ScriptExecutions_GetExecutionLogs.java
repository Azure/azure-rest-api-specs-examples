import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.ScriptOutputStreamType;
import java.util.Arrays;

/** Samples for ScriptExecutions GetExecutionLogs. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_GetExecutionLogs.json
     */
    /**
     * Sample code: ScriptExecutions_GetExecutionLogs.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsGetExecutionLogs(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .scriptExecutions()
            .getExecutionLogsWithResponse(
                "group1",
                "cloud1",
                "addSsoServer",
                Arrays
                    .asList(
                        ScriptOutputStreamType.INFORMATION,
                        ScriptOutputStreamType.fromString("Warnings"),
                        ScriptOutputStreamType.fromString("Errors"),
                        ScriptOutputStreamType.OUTPUT),
                Context.NONE);
    }
}
