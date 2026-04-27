
import com.azure.resourcemanager.azurestackhci.models.ArcConnectivityProperties;
import com.azure.resourcemanager.azurestackhci.models.ArcSetting;
import com.azure.resourcemanager.azurestackhci.models.ServiceConfiguration;
import com.azure.resourcemanager.azurestackhci.models.ServiceName;
import java.util.Arrays;

/**
 * Samples for ArcSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/PatchArcSetting.json
     */
    /**
     * Sample code: Patch ArcSetting.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void patchArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        ArcSetting resource = manager.arcSettings()
            .getWithResponse("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withConnectivityProperties(new ArcConnectivityProperties().withEnabled(true).withServiceConfigurations(
                Arrays.asList(new ServiceConfiguration().withServiceName(ServiceName.WAC).withPort(6516L))))
            .apply();
    }
}
