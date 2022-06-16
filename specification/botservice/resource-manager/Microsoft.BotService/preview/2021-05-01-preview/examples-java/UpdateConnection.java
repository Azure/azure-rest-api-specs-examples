import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.models.ConnectionSetting;
import com.azure.resourcemanager.botservice.models.ConnectionSettingParameter;
import com.azure.resourcemanager.botservice.models.ConnectionSettingProperties;
import java.util.Arrays;

/** Samples for BotConnection Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateConnection.json
     */
    /**
     * Sample code: Update Connection Setting.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateConnectionSetting(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        ConnectionSetting resource =
            manager
                .botConnections()
                .getWithResponse("OneResourceGroupName", "samplebotname", "sampleConnection", Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new ConnectionSettingProperties()
                    .withClientId("sampleclientid")
                    .withClientSecret("samplesecret")
                    .withScopes("samplescope")
                    .withServiceProviderId("serviceproviderid")
                    .withServiceProviderDisplayName("serviceProviderDisplayName")
                    .withParameters(
                        Arrays
                            .asList(
                                new ConnectionSettingParameter().withKey("key1").withValue("value1"),
                                new ConnectionSettingParameter().withKey("key2").withValue("value2"))))
            .withEtag("etag1")
            .apply();
    }
}
