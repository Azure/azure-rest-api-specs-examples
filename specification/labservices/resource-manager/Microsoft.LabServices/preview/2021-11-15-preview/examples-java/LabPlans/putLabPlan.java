import com.azure.resourcemanager.labservices.models.AutoShutdownProfile;
import com.azure.resourcemanager.labservices.models.ConnectionProfile;
import com.azure.resourcemanager.labservices.models.ConnectionType;
import com.azure.resourcemanager.labservices.models.EnableState;
import com.azure.resourcemanager.labservices.models.LabPlanNetworkProfile;
import com.azure.resourcemanager.labservices.models.ShutdownOnIdleMode;
import com.azure.resourcemanager.labservices.models.SupportInfo;
import java.time.Duration;

/** Samples for LabPlans CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/putLabPlan.json
     */
    /**
     * Sample code: putLabPlan.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void putLabPlan(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .labPlans()
            .define("testlabplan")
            .withRegion("westus")
            .withExistingResourceGroup("testrg123")
            .withDefaultConnectionProfile(
                new ConnectionProfile()
                    .withWebSshAccess(ConnectionType.NONE)
                    .withWebRdpAccess(ConnectionType.NONE)
                    .withClientSshAccess(ConnectionType.PUBLIC)
                    .withClientRdpAccess(ConnectionType.PUBLIC))
            .withDefaultAutoShutdownProfile(
                new AutoShutdownProfile()
                    .withShutdownOnDisconnect(EnableState.ENABLED)
                    .withShutdownWhenNotConnected(EnableState.ENABLED)
                    .withShutdownOnIdle(ShutdownOnIdleMode.USER_ABSENCE)
                    .withDisconnectDelay(Duration.parse("00:05"))
                    .withNoConnectDelay(Duration.parse("01:00"))
                    .withIdleDelay(Duration.parse("01:00")))
            .withDefaultNetworkProfile(
                new LabPlanNetworkProfile()
                    .withSubnetId(
                        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"))
            .withSharedGalleryId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig")
            .withSupportInfo(
                new SupportInfo()
                    .withUrl("help.contoso.com")
                    .withEmail("help@contoso.com")
                    .withPhone("+1-202-555-0123")
                    .withInstructions("Contact support for help."))
            .create();
    }
}
