
import com.azure.resourcemanager.networkanalytics.fluent.models.RoleAssignmentDetailInner;
import com.azure.resourcemanager.networkanalytics.models.DataProductUserRole;
import java.util.Arrays;

/**
 * Samples for DataProducts RemoveUserRole.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/
     * DataProducts_RemoveUserRole_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataProducts_RemoveUserRole_MaximumSet_Gen.
     * 
     * @param manager Entry point to NetworkAnalyticsManager.
     */
    public static void dataProductsRemoveUserRoleMaximumSetGen(
        com.azure.resourcemanager.networkanalytics.NetworkAnalyticsManager manager) {
        manager.dataProducts().removeUserRoleWithResponse("aoiresourceGroupName", "dataproduct01",
            new RoleAssignmentDetailInner().withRoleId("00000000-0000-0000-0000-00000000000")
                .withPrincipalId("00000000-0000-0000-0000-00000000000").withUsername("UserName")
                .withDataTypeScope(Arrays.asList("scope")).withPrincipalType("User")
                .withRole(DataProductUserRole.READER).withRoleAssignmentId("00000000-0000-0000-0000-00000000000"),
            com.azure.core.util.Context.NONE);
    }
}
