
import com.azure.resourcemanager.avs.models.ScriptOutputStreamType;
import java.util.Arrays;

/**
 * Samples for ScriptExecutions GetExecutionLogs.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ScriptExecutions_GetExecutionLogs.json
     */
    /**
     * Sample code: ScriptExecutions_GetExecutionLogs.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsGetExecutionLogs(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().getExecutionLogsWithResponse("group1", "cloud1", "addSsoServer",
            Arrays.asList(ScriptOutputStreamType.INFORMATION, ScriptOutputStreamType.fromString("Warnings"),
                ScriptOutputStreamType.fromString("Errors"), ScriptOutputStreamType.OUTPUT),
            com.azure.core.util.Context.NONE);
    }
}
