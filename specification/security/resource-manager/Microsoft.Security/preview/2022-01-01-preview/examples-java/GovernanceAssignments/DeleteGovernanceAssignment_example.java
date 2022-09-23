import com.azure.core.util.Context;

/** Samples for GovernanceAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceAssignments/DeleteGovernanceAssignment_example.json
     */
    /**
     * Sample code: Delete security assignment.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteSecurityAssignment(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .governanceAssignments()
            .deleteWithResponse(
                "subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012",
                "6b9421dd-5555-2251-9b3d-2be58e2f82cd",
                "6634ff9f-127b-4bf2-8e6e-b1737f5e789c",
                Context.NONE);
    }
}
