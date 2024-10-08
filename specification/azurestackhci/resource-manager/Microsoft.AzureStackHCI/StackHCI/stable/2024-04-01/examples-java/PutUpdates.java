
import com.azure.resourcemanager.azurestackhci.models.AvailabilityType;
import com.azure.resourcemanager.azurestackhci.models.HciUpdate;
import com.azure.resourcemanager.azurestackhci.models.State;
import com.azure.resourcemanager.azurestackhci.models.UpdatePrerequisite;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for Updates Put.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * PutUpdates.json
     */
    /**
     * Sample code: Put a specific update.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void putASpecificUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        HciUpdate resource = manager.updates()
            .getWithResponse("testrg", "testcluster", "Microsoft4.2203.2.32", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withInstalledDate(OffsetDateTime.parse("2022-04-06T14:08:18.254Z"))
            .withDescription("AzS Update 4.2203.2.32").withState(State.INSTALLED)
            .withPrerequisites(Arrays.asList(new UpdatePrerequisite().withUpdateType("update type")
                .withVersion("prerequisite version").withPackageName("update package name")))
            .withPackagePath("\\\\SU1FileServer\\SU1_Infrastructure_2\\Updates\\Packages\\Microsoft4.2203.2.32")
            .withPackageSizeInMb(18858.0F).withDisplayName("AzS Update - 4.2203.2.32").withVersion("4.2203.2.32")
            .withPublisher("Microsoft")
            .withReleaseLink("https://docs.microsoft.com/azure-stack/operator/release-notes?view=azs-2203")
            .withAvailabilityType(AvailabilityType.LOCAL).withPackageType("Infrastructure")
            .withAdditionalProperties("additional properties").withProgressPercentage(0.0F)
            .withNotifyMessage("Brief message with instructions for updates of AvailabilityType Notify").apply();
    }
}
