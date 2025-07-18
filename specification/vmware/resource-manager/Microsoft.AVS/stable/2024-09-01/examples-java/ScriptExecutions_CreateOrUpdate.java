
import com.azure.resourcemanager.avs.models.ScriptSecureStringExecutionParameter;
import com.azure.resourcemanager.avs.models.ScriptStringExecutionParameter;
import java.util.Arrays;

/**
 * Samples for ScriptExecutions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ScriptExecutions_CreateOrUpdate.json
     */
    /**
     * Sample code: ScriptExecutions_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void scriptExecutionsCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptExecutions().define("addSsoServer").withExistingPrivateCloud("group1", "cloud1")
            .withScriptCmdletId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource")
            .withParameters(Arrays.asList(
                new ScriptStringExecutionParameter().withName("DomainName").withValue("placeholderDomain.local"),
                new ScriptStringExecutionParameter().withName("BaseUserDN")
                    .withValue("DC=placeholder, DC=placeholder")))
            .withHiddenParameters(Arrays.asList(
                new ScriptSecureStringExecutionParameter().withName("Password").withSecureValue("PlaceholderPassword")))
            .withTimeout("P0Y0M0DT0H60M60S").withRetention("P0Y0M60DT0H60M60S").create();
    }
}
