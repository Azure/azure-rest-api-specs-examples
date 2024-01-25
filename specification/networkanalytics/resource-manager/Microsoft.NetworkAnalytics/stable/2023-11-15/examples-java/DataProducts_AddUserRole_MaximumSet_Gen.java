
import com.azure.resourcemanager.networkanalytics.models.DataProductUserRole;
import com.azure.resourcemanager.networkanalytics.models.RoleAssignmentCommonProperties;
import java.util.Arrays;

/**
 * Samples for DataProducts AddUserRole.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_AddUserRole_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_AddUserRole_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsAddUserRoleMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProducts().addUserRoleWithResponse("aoiresourceGroupName", "dataproduct01",
            new RoleAssignmentCommonProperties().withRoleId("00000000-0000-0000-0000-00000000000")
                .withPrincipalId("00000000-0000-0000-0000-00000000000").withUsername("UserName")
                .withDataTypeScope(Arrays.asList("scope")).withPrincipalType("User")
                .withRole(DataProductUserRole.READER),
            com.azure.core.util.Context.NONE);
    }
}
