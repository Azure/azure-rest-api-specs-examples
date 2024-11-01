
import com.azure.resourcemanager.streamanalytics.models.AuthenticationMode;
import com.azure.resourcemanager.streamanalytics.models.EventGridEventSchemaType;
import com.azure.resourcemanager.streamanalytics.models.EventGridStreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.EventHubV2StreamInputDataSource;
import com.azure.resourcemanager.streamanalytics.models.StorageAccount;
import com.azure.resourcemanager.streamanalytics.models.StreamInputProperties;
import java.util.Arrays;

/**
 * Samples for Inputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Create_EventGrid.json
     */
    /**
     * Sample code: Create an Event Grid input.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void
        createAnEventGridInput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().define("input7970").withExistingStreamingjob("sjrg3467", "sj9742")
            .withProperties(new StreamInputProperties().withDatasource(new EventGridStreamInputDataSource()
                .withSubscriber(new EventHubV2StreamInputDataSource().withConsumerGroupName("sdkconsumergroup")
                    .withEventHubName("sdkeventhub").withPartitionCount(16).withServiceBusNamespace("sdktest")
                    .withSharedAccessPolicyName("RootManageSharedAccessKey")
                    .withSharedAccessPolicyKey("fakeTokenPlaceholder").withAuthenticationMode(AuthenticationMode.MSI))
                .withSchema(EventGridEventSchemaType.CLOUD_EVENT_SCHEMA)
                .withStorageAccounts(Arrays.asList(new StorageAccount().withAccountName("myaccount")
                    .withAccountKey("fakeTokenPlaceholder").withAuthenticationMode(AuthenticationMode.MSI)))
                .withEventTypes(Arrays.asList("Microsoft.Storage.BlobCreated"))))
            .create();
    }
}
