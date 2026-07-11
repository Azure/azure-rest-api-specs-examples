
import com.azure.resourcemanager.azurestackhci.models.ExtensionUpgradeParameters;

/**
 * Samples for Extensions Upgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/Extensions_Upgrade.json
     */
    /**
     * Sample code: Upgrade Machine Extensions.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void upgradeMachineExtensions(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.extensions().upgrade("test-rg", "myCluster", "default", "MicrosoftMonitoringAgent",
            new ExtensionUpgradeParameters().withTargetVersion("1.0.18062.0"), com.azure.core.util.Context.NONE);
    }
}
