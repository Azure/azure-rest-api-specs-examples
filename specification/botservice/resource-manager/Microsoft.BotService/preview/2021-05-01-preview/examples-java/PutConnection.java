import com.azure.resourcemanager.botservice.models.ConnectionSettingParameter;
import com.azure.resourcemanager.botservice.models.ConnectionSettingProperties;
import java.util.Arrays;

/** Samples for BotConnection Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutConnection.json
     */
    /**
     * Sample code: Create Connection Setting.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createConnectionSetting(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .botConnections()
            .define("sampleConnection")
            .withRegion("West US")
            .withExistingBotService("OneResourceGroupName", "samplebotname")
            .withProperties(
                new ConnectionSettingProperties()
                    .withClientId("sampleclientid")
                    .withClientSecret("samplesecret")
                    .withScopes("samplescope")
                    .withServiceProviderId("serviceproviderid")
                    .withParameters(
                        Arrays
                            .asList(
                                new ConnectionSettingParameter().withKey("key1").withValue("value1"),
                                new ConnectionSettingParameter().withKey("key2").withValue("value2"))))
            .withEtag("etag1")
            .create();
    }
}
